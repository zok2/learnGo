async function fetchAndExtractLinks() {
    var targetURL = 'https://www.aiguoba.com/manhua/jiangsha/';
    extractedURLs =[];
    try {
        var response = await fetch(targetURL);
        var data = await response.text();

        var parser = new DOMParser();
        var doc = parser.parseFromString(data, 'text/html');
        var targetDiv = doc.querySelector('.plist.pnormal');

        if (targetDiv) {
            var links = targetDiv.querySelectorAll('a');
            links.forEach(link => {
                var linkURL = link.getAttribute('href');
                extractedURLs.push(linkURL)
                console.log(linkURL);
            });
        } else {
            console.log('Target div not found.');
        }
    } catch (error) {
        console.error('An error occurred:', error);
    }
}

async function fetchAndDownloadImages(url) {
    try {
        const response = await fetch(url);
        const pageContent = await response.text();

        // 步骤 3: 解析页面内容以获取图片信息
        const parser = new DOMParser();
        const doc = parser.parseFromString(pageContent, 'text/html');
        const imagesDiv = doc.getElementById('imagesOld');
        if (imagesDiv) {
            const images = imagesDiv.querySelectorAll('img');
            images.forEach((image, index) => {
                const imageURL = image.getAttribute('data-original');
                const imageName = image.getAttribute('alt');

                // 创建下载链接
                const downloadLink = document.createElement('a');
                downloadLink.href = imageURL;
                downloadLink.download = `${imageName}_${index + 1}.jpg`;
                downloadLink.textContent = `Download ${imageName}_${index + 1}.jpg`;
                document.body.appendChild(downloadLink);
            });
        } else {
            console.log('Images div not found.');
        }
    } catch (error) {
        console.error('An error occurred:', error);
    }
}
extractedURLs.forEach(url => {
    fetchAndDownloadImages(url);
    setTimeout()
});
// 循环遍历URL数组并请求页面内容