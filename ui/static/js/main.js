document.addEventListener('DOMContentLoaded', function () {
  const links = document.querySelectorAll('.nav-container a');

  function updateActiveLink() {
    const currentPath = window.location.pathname;
    links.forEach(link => {
      if (link.getAttribute('href') === currentPath) {
        link.classList.add('active');
      } else {
        link.classList.remove('active');
      }
    });
  }

  updateActiveLink();
  const articleText = document.querySelector('blog-post').innerText;
  const readTime = document.querySelector('read-time');
  const timeToRead = calculateReadTime(articleText);
  readTime.textContent = `${timeToRead} mins`
});

function calculateReadTime(text) {
  const words = text.split(/\s+/g);
  const wordCount = words.length;

  // Assume an average reading speed of 200 words per minute
  const readingSpeed = 200;

  const readTime = Math.ceil(wordCount / readingSpeed);

  return readTime;
}
