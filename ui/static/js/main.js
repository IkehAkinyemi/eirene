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
});
