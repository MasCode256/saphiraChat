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
      "// Кол-во контактов" + contacts_count;
  } catch (error) {
    console.error("Ошибка при получении контактов:", error);
  }
});
