import { useState } from 'react';
import setRenderTargets from './reactRenderer';

type CounterProps = {
    startCount: number,
};

const Counter = ({ props }: { props: CounterProps }) => {
    let [count, setCount] = useState(props.startCount);

    let changer = (amount: number) => {
        return () => setCount(count + amount);
    }

    let incr = changer(1);
    let decr = changer(-1);

    return <div className="mb-2 space-x-2">
        <button onClick={decr} className="px-2 py-1 cursor-pointer border rounded">-</button>
        <span>{count}</span>
        <button onClick={incr} className="px-2 py-1 cursor-pointer border rounded">+</button>
    </div>;
};

setRenderTargets('counter', Counter);
