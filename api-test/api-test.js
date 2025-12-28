const url = "http://localhost:8000/users"; // replace with your API
const totalRequests = 1000;
const batchSize = 100; // number of requests at a time

async function makeRequest() {
  try {
    const response = await fetch(url);
    const data = await response.json(); // if API returns JSON
    return data;
  } catch (error) {
    console.error("Request failed:", error);
    return null;
  }
}
async function runLoadTest() {
  let completedRequests = 0;

  while (completedRequests < totalRequests) {
    const requests = [];
    for (let i = 0; i < batchSize && completedRequests < totalRequests; i++) {
      requests.push(makeRequest());
      completedRequests++;
    }
    await Promise.all(requests);
    console.log(`Completed ${completedRequests} / ${totalRequests} requests`);
  }
  console.log("All requests completed");
}

runLoadTest();