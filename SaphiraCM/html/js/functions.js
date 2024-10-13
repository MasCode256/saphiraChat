document.addEventListener("DOMContentLoaded", async function () {
  console.log("Functions is ready!");
});

var contacts_count = 0;

var names = [];
var ips = [];
var colors = [];
var keys = [];

async function fetchTextFromUrl(url) {
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    const text = await response.text();
    return text;
  } catch (error) {
    console.error("Ошибка при получении данных:", error);
    throw error; // Проброс ошибки, чтобы она могла быть обработана вызывающим кодом
  }
}
