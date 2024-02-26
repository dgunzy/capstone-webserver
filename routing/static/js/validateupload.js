document.addEventListener('DOMContentLoaded', function () {
    const fileInput = document.getElementById('dropzone-file');
    const submitButton = document.getElementById('submit-button'); 
    const notificationDiv = document.getElementById('notification');

    fileInput.addEventListener('change', function(event) {
        const file = event.target.files[0];
        if (!file || file.type !== "application/pdf") {
            alert("Please select a PDF file.");
            return;
        }

        const fileReader = new FileReader();
        
        fileReader.onload = function() {
            const typedArray = new Uint8Array(this.result);
            pdfjsLib.getDocument(typedArray).promise.then(function(pdf) {
                let pdfText = '';
                const queue = [];
                for (let pageNum = 1; pageNum <= pdf.numPages; pageNum++) {
                    queue.push(pdf.getPage(pageNum).then(function(page) {
                        return page.getTextContent().then(function(textContent) {
                            return textContent.items.map(item => item.str).join(' ');
                        });
                    }));
                }
                
                Promise.all(queue).then(function(texts) {
                    pdfText = texts.join('\n');
                    pdfText = pdfText.replace(/\s+/g, ' ');
                    pdfText = pdfText.replace(/\.{2,}/g, ' ');
                    pdfText = pdfText.replace(/\S*@\S*/g, '');
                    notificationDiv.innerText = 'PDF processing complete. You can now submit the text.';
                    document.getElementById('extracted-text').value = pdfText; 
                    submitButton.disabled = false; 
                });
            });
        };
        
        fileReader.readAsArrayBuffer(file);
    });

});