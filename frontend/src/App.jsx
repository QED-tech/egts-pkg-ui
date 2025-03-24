import "./App.css";
import React, { useState } from "react";
import { decodeEgts } from "./services/egts";
import ViewParsedPgk from "./components/ViewParsedPkg";

function App() {
  const [hexString, setHexString] = useState("");
  const [parsedPackage, setParsedPackage] = useState(null);

  const handleParse = async () => {
    const parsedEgts = await decodeEgts(hexString);
    console.log(parsedEgts.data);
    setParsedPackage(parsedEgts.data);
  };

  return (
    <div className="app">
      <div className="app-header">
        <div class="text-center text-5xl font-extrabold leading-none tracking-tight m-8">
          <span class="text-gray-900">EGTS Debug</span>
        </div>

        <div className="m-8">
          <pre>
            Пример пакета:
            0100020b002300020001871800010011e70300000202101500b6739d1b4fba3a9ed227bc350000000000000000003b07
          </pre>
        </div>
        <input
          className="hex-input"
          value={hexString}
          onChange={(e) => setHexString(e.target.value)}
          placeholder="EGTS Package hex string"
        />

        <button className="text-white" onClick={handleParse}>
          Decode
        </button>
      </div>

      {parsedPackage && (
        <ViewParsedPgk parsedPackage={parsedPackage} hexString={hexString} />
      )}
    </div>
  );
}

export default App;
