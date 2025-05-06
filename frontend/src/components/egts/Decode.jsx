import JsonView from "@uiw/react-json-view";
import { monokaiTheme } from "@uiw/react-json-view/monokai";
import {
  Tooltip,
  TooltipTrigger,
  TooltipContent,
} from "@/components/ui/tooltip";
import { decodeEgts } from "@/services/egts";
import React, { useState, useEffect } from "react";
import Header from "@/components/ui/header.js";

const formatHexString = (hexString) => {
  if (hexString == "") {
    return "";
  }
  return hexString.match(/.{1,2}/g).join("\n");
};

const ViewParsedPgk = () => {
  const [parsedPackage, setParsedPackage] = useState(null);

  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.search);
    const hexPkg = urlParams.get("pkg");

    const fetchData = async () => {
      try {
        const response = await decodeEgts(hexPkg);
        setParsedPackage(response.data);
      } catch (err) {
        console.error("Ошибка при разборе пакета:", err);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="app w-screen">
      <Header></Header>

      <div className="container mx-auto">
        <table className="border border-gray-300 table-auto">
          <thead>
            <tr className="border">
              <th className="px-4 py-2 text-left border">Поле</th>
              <th className="px-4 py-2 text-left border">Hex</th>
              <th className="px-4 py-2 text-left border">Значение</th>
              <th className="px-4 py-2 text-left border">Размер (байты)</th>
            </tr>
          </thead>
          <tbody>
            {parsedPackage &&
              Object.entries(parsedPackage).map(([key, value], index) => {
                return (
                  <tr key={index} className="border-b hover:bg-gray-50">
                    <td className="px-4 py-2 border align-top">
                      <Tooltip key={index}>
                        <TooltipTrigger asChild>
                          <span className="hexField cursor-pointer text-blue-600 font-bold">
                            {key}
                          </span>
                        </TooltipTrigger>
                        <TooltipContent className="overflow-scroll">
                          <strong>Описание:</strong>{" "}
                          <pre>{value.description}</pre>
                        </TooltipContent>
                      </Tooltip>
                    </td>
                    <td className="px-4 py-2 border align-top">
                      <span>{formatHexString(value.hex)}</span>
                    </td>
                    <td className="px-4 py-2 border align-top">
                      {key === "SFRD" ? (
                        <JsonView
                          value={JSON.parse(value.value)}
                          theme={monokaiTheme}
                          displayDataTypes={false}
                          displayObjectSize={false}
                          enableClipboard={true}
                        />
                      ) : (
                        value.value
                      )}
                    </td>
                    <td className="px-4 py-2 border align-top">
                      {value.bytes_size}
                    </td>
                  </tr>
                );
              })}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ViewParsedPgk;
