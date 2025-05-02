import "./App.css";
import React, { useState, useEffect, useRef } from "react";
import { Link } from "react-router";
import History from "./components/egts/History.jsx";
import Header from "./components/ui/header.js";

function App() {
  const [hexString, setHexString] = useState("");
  const textAreaRef = useRef(null);

  useEffect(() => {
    if (textAreaRef.current) {
      textAreaRef.current.style.height =
        textAreaRef.current.scrollHeight + "px";
    }
  }, [hexString]);

  return (
    <div className="app">
      <Header></Header>
      <div className="app-header">
        <div className="text-center text-5xl font-extrabold leading-none tracking-tight m-8">
          <span className="text-gray-700">Decode</span>
        </div>

        <div className="m-8">
          <pre>
            Пример пакета EGTS:
            0100020b002300020001871800010011e70300000202101500b6739d1b4fba3a9ed227bc350000000000000000003b07
          </pre>
        </div>
        <textarea
          className="h-auto w-full p-3 border border-gray-300 rounded-lg focus:ring focus:ring-blue-300 outline-none text-sm resize-none overflow-hidden"
          value={hexString}
          ref={textAreaRef}
          onChange={(e) => setHexString(e.target.value)}
          onInput={(e) => {
            e.target.style.height = e.target.scrollHeight + "px";
          }}
          placeholder="Введите EGTS пакет в HEX формате"
        />

        <Link to={{ pathname: "/hex", search: `?pkg=${hexString}` }}>
          <button className="text-white">Decode</button>
        </Link>
      </div>

      <History />
    </div>
  );
}

export default App;
