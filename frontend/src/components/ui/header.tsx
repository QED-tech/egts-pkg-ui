import React from "react";
import { Link } from "react-router";

const Header = () => {
  return (
    <div className="header text-gray-900 p-4">
      <nav className="flex justify-center space-x-8">
        <Link to="/" className="px-4 py-2 text-3xl">
          Decode
        </Link>
        <Link to="/encode" className="px-4 py-2 rounded text-3xl">
          Encode
        </Link>
      </nav>
    </div>
  );
};


export default Header