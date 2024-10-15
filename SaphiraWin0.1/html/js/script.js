document.addEventListener("DOMContentLoaded", async function () {
  console.log("Script is ready!");

  try {
    const contactsText = await fetchTextFromUrl(
      "http://localhost:1110?get=contacts"
    );
    const contactsCount = parseInt(contactsText, 10);
    console.log("CC =", contactsCount);
    contacts_count = contactsCount;

    document.getElementById("cc").innerText =
      "// Кол-во контактов: " + contacts_count;
  } catch (error) {
    console.error("Ошибка при получении контактов:", error);
  }

  const contacts = document.getElementById("contacts");

  for (let index = 0; index < contacts_count; index++) {
    var name = "NULL";
    var ip = "NULL";
    var color = "NULL";
    var key = "NULL";

    try {
      const text = await fetchTextFromUrl(
        "http://localhost:1110?name=" + index
      );

      name = text;
      console.log("name: " + name);
    } catch (error) {
      console.error("Ошибка при получении имени:", error);
    }

    try {
      const text = await fetchTextFromUrl("http://localhost:1110?ip=" + index);

      ip = text.substring(0, text.length - 1);
    } catch (error) {
      console.error("Ошибка при получении ip:", error);
    }

    try {
      const text = await fetchTextFromUrl(
        "http://localhost:1110?color=" + index
      );

      color = text;
    } catch (error) {
      console.error("Ошибка при получении цвета:", error);
    }

    names.push(name);
    ips.push(ip);
    colors.push(color);
    keys.push(key);

    contacts.innerHTML +=
      "<p></p><button onclick='switch_contact(" +
      index +
      "); load_settings(current)' class='btn c2' style='--c2:" +
      color +
      "'>" +
      name +
      "(" +
      ip +
      ")" +
      "</button>";
  }

  switch_contact(0);
});
