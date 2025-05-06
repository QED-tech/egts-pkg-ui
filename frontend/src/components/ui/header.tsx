import React from "react";
import { NavLink } from "react-router";

const Header = () => {
  return (
    <header className="bg-gray-800 py-1 text-white">
      <nav className="container mx-auto flex justify-center space-x-8">
        <NavLink
          to="/"
          className={({ isActive }) =>
            `px-4 py-2 text-2xl font-semibold transition duration-300 ${
              isActive
                ? "text-blue-600 border-b-2 border-blue-600"
                : "text-gray-700 hover:text-blue-500"
            }`
          }
        >
          Decode
        </NavLink>
      </nav>
    </header>
  );
};

export default Header;
