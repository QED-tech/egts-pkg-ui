import { useState, useEffect } from "react";
import { fetchHistory } from "@/services/egts";
import { Link } from "react-router";

const History = () => {
  const [history, setHistory] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetch = async () => {
      setLoading(true);
      setError(null);
      try {
        const response = await fetchHistory();
        setHistory(response.data);
      } catch (err) {
        console.error("Ошибка при получении истории запросов:", err);
        setError("Не удалось загрузить историю запросов.");
      } finally {
        setLoading(false);
      }
    };

    fetch();
  }, []);

  if (loading) {
    return (
      <div className="text-lg text-gray-600 italic">
        Загрузка истории запросов...
      </div>
    );
  }

  if (error) {
    return <div className="text-red-500 text-sm">{error}</div>;
  }

  if (!history.length) {
    return (
      <div className="mb-4 text-xl font-semibold text-gray-700">
        Нет истории запросов
      </div>
    );
  }

  return (
    <div className="bg-white p-6 mt-8">
      <h2 className="text-xl font-semibold text-gray-800 mb-4">
        История запросов:
      </h2>
      <ul>
        {history.map((historyItem) => (
          <li
            key={historyItem.id}
            className="mb-2 last:mb-0 border-b border-gray-200 pb-2 last:border-b-0"
          >
            <Link
              to={{ pathname: "/hex", search: `?pkg=${historyItem.hex}` }}
              className="block hover:bg-gray-100 rounded-md p-2 transition duration-150 ease-in-out"
            >
              <p className="text-sm text-gray-700 break-words">
                {truncate(historyItem.hex, 150)}
              </p>
              {historyItem.timestamp && (
                <p className="text-xs text-gray-500 mt-1">
                  {new Date(historyItem.timestamp).toLocaleString()}
                </p>
              )}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

function truncate(str, n) {
  return str.length > n ? str.slice(0, n - 1) + "....." : str;
}

export default History;
