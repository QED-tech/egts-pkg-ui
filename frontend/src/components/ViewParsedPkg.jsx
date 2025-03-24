import JsonView from "@uiw/react-json-view";
import { monokaiTheme } from "@uiw/react-json-view/monokai";

const formatHexString = (hexString) => {
  return hexString.match(/.{1,20}/g).join("\n");
};

const ViewParsedPgk = ({ parsedPackage, hexString }) => {
  let offset = 0;

  return (
    <div className="overflow-x-auto mt-5">
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
          {Object.entries(parsedPackage).map(([key, value], index) => {
            const hexField = hexString.slice(
              offset,
              offset + value.bytes_size * 2
            );
            offset += value.bytes_size * 2;

            return (
              <tr key={index} className="border-b hover:bg-gray-50">
                <td className="px-4 py-2 font-bold border"><pre>{key}</pre></td>
                <td className="px-4 py-2 font-bold border">
                  <span>{formatHexString(hexField)}</span>
                </td>
                <td className="px-4 py-2 border">
                  {key === "SFRD" ? (
                    <JsonView
                      value={JSON.parse(value.value)}
                      theme={monokaiTheme}
                      displayDataTypes={false}
                      displayObjectSize={false}
                      enableClipboard={false}
                    />
                  ) : value.value}
                </td>
                <td className="px-4 py-2 border">{value.bytes_size}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};

export default ViewParsedPgk;
