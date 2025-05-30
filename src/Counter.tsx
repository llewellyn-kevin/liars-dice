import { useEffect, useState } from 'react';
import setRenderTargets from './reactRenderer';
import type { CounterProps } from './frontend_messages';

interface WsMessage<T> {
    message: string;
    data: T;
};

const Counter = ({ props }: { props: CounterProps }) => {
    let [count, setCount] = useState(props.startCount);
    let [socket, setSocket] = useState<WebSocket | null>(null)

    let changer = (amount: number) => {
        return () => setCount(count + amount);
    }

    let incr = changer(1);
    let decr = changer(-1);

    useEffect(() => {
        const ws = new WebSocket("ws://localhost:8080/ws/counter");
        ws.addEventListener("message", (event: MessageEvent) => {
            const msg: CounterProps = JSON.parse(event.data).data;
            setCount(msg.startCount);
        });
        setSocket(ws);

        return () => {
            ws?.close()
        }
    }, []);

    useEffect(() => {
        if (socket === null) {
            return;
        }

        const message: WsMessage<CounterProps> = {
            message: 'update-counter',
            data: {
                startCount: count,
            },
        };
        socket.send(JSON.stringify(message));
    }, [count])

    // Connection opened
    // useEffect(() => {
    //     socket.addEventListener("open", () => {
    //         socket.send(JSON.stringify({
    //             startCount: count,
    //         }));
    //     });
    // }, [])

    // Listen for messages

    return <div className="mb-2 space-x-2">
        <button onClick={decr} className="px-2 py-1 cursor-pointer border rounded">-</button>
        <span>{count}</span>
        <button onClick={incr} className="px-2 py-1 cursor-pointer border rounded">+</button>
    </div>;
};

setRenderTargets('counter', Counter);
