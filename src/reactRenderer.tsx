import type { JSX } from 'react';
import { createRoot } from 'react-dom/client';

export type ComponentProps<T> = {
    props: T,
};

const setRenderTargets = <T,>(name: string, Foo: (props: ComponentProps<T>) => JSX.Element) => {
    const nodes = document.querySelectorAll<HTMLElement>(`[rt-component="${name}"]`);
    nodes.forEach((node) => {
        const props = JSON.parse(node.dataset.rt ?? '{}');
        createRoot(node).render(<Foo props={props} />);
    });
};

export default setRenderTargets;
