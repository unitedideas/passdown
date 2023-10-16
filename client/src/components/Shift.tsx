// src/components/Shift.tsx
import React from 'react';

type ShiftProps = {
    id: number;
    name: string;
    shift: string;
};

const Shift: React.FC<ShiftProps> = ({ id, name, shift }) => {
    return (
        <tr>
            <td>{id}</td>
            <td>{name}</td>
            <td>{shift}</td>
        </tr>
    );
};

export default Shift;
