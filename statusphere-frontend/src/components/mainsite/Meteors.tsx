'use client'
import clsx from "clsx";
import React from "react";

export const Meteors = ({ number }: { number?: number }) => {
    const meteors = new Array(number || 20).fill(true);
    return (
        <div className={"w-full h-full absolute z-0"}>
            {meteors.map((el, idx) => (
                <span
                    key={"meteor" + idx}
                    className={clsx(
                        "animate-meteor-effect absolute top-1/2 left-1/2 h-0.5 w-0.5 rounded-[9999px] bg-slate-500 shadow-[0_0_0_1px_#ffffff10] rotate-[215deg]",
                        "before:content-[''] before:absolute before:top-1/2 before:transform before:-translate-y-[50%] before:w-[50px] before:h-[1px] before:bg-gradient-to-r before:from-[#64748b] before:to-transparent z-0"
                    )}
                    style={{
                        top: 0,
                        left: Math.floor(Math.random() * 100 / 2) + "vw",
                        animationDelay: Math.random() * 0.6 + 0.2 + "s",
                        animationDuration: Math.floor(Math.random() * 16 + 2) + "s",
                    }}
                ></span>
            ))}
        </div>
    );
};