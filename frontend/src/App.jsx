import "./App.css";
import React, { useState, useEffect, useRef } from "react";
import { Link } from "react-router";
import History from "./components/egts/History.jsx";
import Header from "./components/ui/header.js";
import Footer from "./components/ui/footer.js";

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
    <div className="w-screen flex flex-col h-screen justify-between">
      <Header></Header>

      <main class="container mx-auto">
        <div className="bg-white rounded-lg p-6">
          <div className="text-center text-3xl font-extrabold text-gray-800 mb-4">
            <span>Decode</span>
          </div>

          <div className="mb-6 text-sm text-gray-600 italic">
            <pre className="whitespace-pre-wrap">
              Пример пакета EGTS:
              0100020b002300020001871800010011e70300000202101500b6739d1b4fba3a9ed227bc350000000000000000003b07
            </pre>
          </div>
          <textarea
            className="w-full p-3 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-200 focus:border-blue-500 outline-none text-sm resize-none overflow-hidden mb-4"
            value={hexString}
            ref={textAreaRef}
            onChange={(e) => setHexString(e.target.value)}
            onInput={(e) => {
              e.target.style.height = e.target.scrollHeight + "px";
            }}
            placeholder="Введите EGTS пакет в HEX формате"
          />

          <Link to={{ pathname: "/hex", search: `?pkg=${hexString}` }}>
            <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-300">
              Decode
            </button>
          </Link>
        </div>

        <History />
      </main>

      <Footer />
    </div>
  );
}

export default App;
