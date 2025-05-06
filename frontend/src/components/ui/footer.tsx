import React from "react";

const Footer = () => {
  return (
    <footer className="w-full bg-gray-800 text-white flex flex-col justify-between">
      <div className="container mx-auto py-1">
        <div className="text-center">
          <p className="text-sm">© {new Date().getFullYear()} EGTS Debugger.</p>
        </div>
        <div className="text-center">
          <a
            href="https://t.me/VladislavBerezovskiy"
            target="_blank"
            rel="noopener noreferrer"
            className="text-blue-400 hover:text-blue-300 transition-colors duration-200 text-sm"
          >
            По вопросам приложения: Telegram @VladislavBerezovskiy
          </a>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
