import React, { useState } from 'react';
import { createRoot } from 'react-dom/client';

function Counter() {
    let [count, setCount] = useState(0);

    let changer = (amount: number) => {
        return () => setCount(count + amount);
    }

    let incr = changer(1);
    let decr = changer(-1);

    return <div>
        <button onClick={decr}>-</button>
        <span>{count}</span>
        <button onClick={incr}>+</button>
    </div>;
};

const domNode = document.getElementById('react-counter');
if (domNode !== null) {
    const root = createRoot(domNode);
    root.render(<Counter />);
}
