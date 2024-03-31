'use client'

import {useEffect, useRef, useState} from "react";
import {Button} from "@/components/ui/button";
import {useRouter} from "next/router";


interface CompanyCountProps {
    count: number;
}

export function CompanyCount(props: CompanyCountProps) {
    const currentCount = useRef(0);
    const [shouldRender, setShouldRender] = useState(false);


    useEffect(() => {
        const interval = setInterval(() => {
            if (currentCount.current < props.count) {
                currentCount.current += 1;
                setShouldRender((prev) => !prev);
            } else {
                clearInterval(interval);
            }
        }, 5);
    }, []);

    return <div className="flex justify-center">
        <div className="text-2xl tracking-tight">
            Service status for <a href={"/statusphere/directory"}><b className="text-[#00243c]">{currentCount.current}</b></a> companies!
        </div>
    </div>

}
