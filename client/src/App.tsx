import React, { useEffect, useState } from 'react';
import ShiftsTable from './components/ShiftsTable';

type ShiftData = {
  id: number;
  name: string;
  shift: string;
};

const App: React.FC = () => {
  const [shifts, setShifts] = useState<ShiftData[]>([]);
  const [showShifts, setShowShifts] = useState<boolean>(window.location.pathname === '/api/v1/shifts');

  useEffect(() => {
    if(showShifts) {
      const fetchData = async () => {
        const response = await fetch('/api/v1/shifts');
        const data = await response.json();
        setShifts(data);
      };

      fetchData();
    }
  }, [showShifts]);

  return (
      <div>
        {showShifts ? (
            <>
              <h1>Shift Dashboard</h1>
              <ShiftsTable shifts={shifts} />
            </>
        ) : (
            <h1>Hello World!</h1>
        )}
      </div>
  );
};

export default App;
