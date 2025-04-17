import { useState, useEffect } from "react";
import { fetchHistory } from "@/services/egts";
import { Link } from "react-router";

const History = () => {
  const [history, setHistory] = useState([]);
  useEffect(() => {
    const fetch = async () => {
      try {
        const response = await fetchHistory();
        setHistory(response.data);
      } catch (err) {
        console.error("Ошибка при получении истории запросов:", err);
      }
    };

    fetch();
  }, []);

  if (!history.length)
    return (
      <div className="text-left text-2xl font-extrabold leading-none tracking-tight">
        <span className="text-gray-900">Нет истории запросов</span>
      </div>
    );

  return (
    <>
      <div className="text-left text-2xl font-extrabold leading-none tracking-tight">
        <span className="text-gray-900">История запросов:</span>
      </div>
      <div className="mt-5">
        {history &&
          history.map((historyItem) => {
            return (
              <Link
                to={{ pathname: "/hex", search: `?pkg=${historyItem.hex}` }}
              >
                <p>{truncate(historyItem.hex, 150)}</p>
              </Link>
            );
          })}
      </div>
    </>
  );
};

function truncate(str, n){
  return (str.length > n) ? str.slice(0, n-1) + '.....' : str;
};

export default History;
