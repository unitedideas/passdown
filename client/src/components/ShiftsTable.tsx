// src/components/ShiftsTable.tsx
import React from 'react';
import Shift from './Shift';

type ShiftData = {
    id: number;
    name: string;
    shift: string;
};

type ShiftsTableProps = {
    shifts: ShiftData[];
};

const ShiftsTable: React.FC<ShiftsTableProps> = ({ shifts }) => {
    return (
        <table>
            <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Shift</th>
            </tr>
            </thead>
            <tbody>
            {shifts.map((shift) => (
                <Shift key={shift.id} {...shift} />
            ))}
            </tbody>
        </table>
    );
};

export default ShiftsTable;
