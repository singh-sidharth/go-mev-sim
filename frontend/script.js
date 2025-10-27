const mempoolTable = document.querySelector("#mempool-table tbody");
const blockTable = document.querySelector("#block-table tbody");
const template = document.querySelector("#tx-row-template"); // HTML template for transaction row

async function fetchData() {
    try {
        // fetch both mempool and block data concurrently
        // helps reduce wait time
        const [mempoolRes, blockRes] = await Promise.all([
            fetch("http://localhost:8080/mempool"),
            fetch("http://localhost:8080/block")
        ]);

        const [mempoolData, blockData] = await Promise.all([
            mempoolRes.json(),
            blockRes.json()
        ]);

        populateTable(mempoolTable, mempoolData);
        populateTable(blockTable, blockData);
    } catch (err) {
        console.error("Error fetching data:", err);
    }
}

function populateTable(tbody, data) {
    // clear existing rows
    tbody.innerHTML = "";

    data.forEach(tx => {
        // clone the template
        const clone = template.content.cloneNode(true);

        // populate fields
        clone.querySelector(".tx-id").textContent = tx.id;
        clone.querySelector(".tx-sender").textContent = tx.sender;
        clone.querySelector(".tx-amount").textContent = tx.amount.toFixed(2);
        clone.querySelector(".tx-profit").textContent = tx.profit.toFixed(2);
        clone.querySelector(".tx-timestamp").textContent = new Date(tx.timestamp).toLocaleTimeString();

        // append to tbody
        tbody.appendChild(clone);
    });
}


// Refresh every second
setInterval(fetchData, 1000);
fetchData();
