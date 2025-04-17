async function decodeEgts(hexString) {
  const response = await fetch(
    `http://localhost:9090/api/v1/egts/decode?package=${hexString}`,
    {
      method: "GET",
    }
  );

  return await response.json();
}

async function fetchHistory() {
  const response = await fetch(`http://localhost:9090/api/v1/egts/history`, {
    method: "GET",
  });

  return await response.json();
}

export { decodeEgts, fetchHistory };
