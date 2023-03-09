import { getObjects } from "./src/client";
import { CatalogObject } from "./src/types";

console.log(`loaded index.js`);

$(async function () {
  console.log("ready");
  try {
    const objects = await getObjects();
    fillObjects(objects);
  } catch (e) {
    showText(`sorry, ${e.message ?? "an error happened"}`);
  }

  console.log("done");
});

function fillObjects(objects: CatalogObject[]) {
  const objectsEl = $("#objects");

  objects.forEach((o, _) => {
    objectsEl.append(`
              <tr>
                  <th scope="row">${o.id}</th>
                  <td>${o.name}</td>
              </tr>
          `);
  });

  $("#objects-table").show();
}

function showText(text: string) {
  $("#delayed-text").text(text);
}
