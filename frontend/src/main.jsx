import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.jsx";
import DecodeHexPkg from "./components/decode/DecodeHexPkg.jsx";
import { Routes, Route, BrowserRouter } from "react-router";

createRoot(document.getElementById("root")).render(
  <BrowserRouter>
    <StrictMode>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/hex" element={<DecodeHexPkg />} />
      </Routes>
    </StrictMode>
  </BrowserRouter>
);
