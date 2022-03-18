const LOCAL_STORAGE_KEY = 'mailhog-theme';
const LOCAL_STORAGE_DATA = JSON.parse(localStorage.getItem(LOCAL_STORAGE_KEY));

const DARK_THEME_PATH = 'css/bootstrap-cyborg.min.css';
const LIGHT_THEME_PATH = 'css/bootstrap-3.3.2.min.css';

const GITHUB_LOGO = document.getElementById('github');
const DARK_STYLE_LINK = document.getElementById('theme-dark');
const LIGHT_STYLE_LINK = document.getElementById('theme-light');
const THEME_TOGGLER = document.getElementById('theme-toggler');

let isDark = LOCAL_STORAGE_DATA && LOCAL_STORAGE_DATA.isDark;
if (isDark) {
  setTheme(isDark)
}

function toggleTheme() {
  isDark = !isDark;

  setTheme(isDark);
  storeTheme(isDark);
}

function storeTheme(isDark) {
  localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify({isDark}));
}

function setTheme(isDark) {
  title = 'Switch beetwen dark and light mode';
  if (isDark) {
    GITHUB_LOGO.setAttribute('class', 'theme-dark');

    DARK_STYLE_LINK.setAttribute('href', DARK_THEME_PATH);
    LIGHT_STYLE_LINK.setAttribute('href', '');
  
    THEME_TOGGLER.innerHTML = '🌙 Dark';
    THEME_TOGGLER.setAttribute('title', title + ' (currently dark mode)');
  } else {
    GITHUB_LOGO.setAttribute('class', '');

    DARK_STYLE_LINK.setAttribute('href', '');
    LIGHT_STYLE_LINK.setAttribute('href', LIGHT_THEME_PATH);

    THEME_TOGGLER.innerHTML = '☀️ Light';
    THEME_TOGGLER.setAttribute('title', title + ' (currently light mode)');
  }
}