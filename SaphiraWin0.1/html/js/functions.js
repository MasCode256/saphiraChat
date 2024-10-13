document.addEventListener("DOMContentLoaded", async function () {
  console.log("Functions is ready!");
});

var is_selected_contact = false;
var current = 0;

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

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function switch_contact(index) {
  is_selected_contact = true;
  console.log("Switching contact: " + index);
  document.getElementById("title").innerHTML =
    "Контакт: " + names[index] + " (" + ips[index] + ")";

  current = index;
}

async function delete_contact(index) {
  console.log("Попытка удаления контакта");

  if (is_selected_contact) {
    fetchTextFromUrl("http://localhost:1110?delete_contact=" + current);
    location.reload();
  } else {
    const output = document.getElementById("output");
    var temp = output.innerHTML;

    output.style = "color: #f00; text-shadow: 0 0 20px #f00, 0 0 10px #f00";
    output.innerHTML =
      ">_ error:11:Аккаунт не выбран! Удаление 'undefined' невозможно.";

    await sleep(3000);

    output.style = "color: #fff; text-shadow: 0 0 20px #fff, 0 0 10px #fff";
    output.innerHTML = temp;
  }
}

var uis = ["chat", "settings"];

async function switch_ui(id) {
  // Получаем элемент по его id
  var element = document.getElementById(id);

  // Удаляем класс 'classToRemove' из списка классов элемента
  for (let index = 0; index < uis.length; index++) {
    const element2 = document.getElementById(uis[index]);
    element2.classList += " invisible";
  }

  element.classList.remove("invisible");
}

async function load_settings(index) {
  try {
    const text = await fetchTextFromUrl("http://localhost:1110?name=" + index);

    data = text.substring(0, text.length - 1);
    document.getElementById("name").value = data;
  } catch (error) {
    console.error("Ошибка при получении ip:", error);
  }

  try {
    const text = await fetchTextFromUrl("http://localhost:1110?ip=" + index);

    data = text.substring(0, text.length - 1);
    document.getElementById("ip").value = data;
  } catch (error) {
    console.error("Ошибка при получении ip:", error);
  }

  try {
    const text = await fetchTextFromUrl("http://localhost:1110?color=" + index);

    data = text.substring(0, text.length - 1);
    document.getElementById("color").value = data;
  } catch (error) {
    console.error("Ошибка при получении ip:", error);
  }

  try {
    const text = await fetchTextFromUrl("http://localhost:1110?key=" + index);

    data = text.substring(0, text.length - 1);
    document.getElementById("key").value = data;
  } catch (error) {
    console.error("Ошибка при получении ip:", error);
  }
}

async function confirm(index) {
  const name = document.getElementById("name").value;
  const ip = document.getElementById("ip").value;
  const color = document.getElementById("color").value;
  const key = document.getElementById("key").value;

  await fetchTextFromUrl(
    "http://localhost:1110?set_name=" + index + "/" + name
  );
  await fetchTextFromUrl("http://localhost:1110?set_ip=" + index + "/" + ip);
  await fetchTextFromUrl(
    "http://localhost:1110?set_color=" +
      index +
      "/" +
      color.substring(1, color.length)
  );
  await fetchTextFromUrl("http://localhost:1110?set_key=" + index + "/" + key);

  location.reload();
}
