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

    return <div>
        <button onClick={decr}>-</button>
        <span>{count}</span>
        <button onClick={incr}>+</button>
    </div>;
};

setRenderTargets('counter', Counter);
