import {useEffect, useRef, useState} from "react";
import axios from "@/utils/axios";
import {StatusPage} from "@/model/StatusPage";


export function CompanyCount() {
    const [count, setCount] = useState(0);
    const currentCount = useRef(0);
    const [shouldRender, setShouldRender] = useState(false);

    useEffect(() => {
        const getCompanyCount = async () => {
            try {
                const response = await axios.get(
                    '/api/v1/statusPages/count'
                );
                setCount(response.data.statusPageCount)
                return response.data.statusPageCount;
            } catch (err) {
                console.log(err);
            }
        };

        getCompanyCount().then((a) => {
            const interval = setInterval(() => {
                if (currentCount.current < a) {
                    currentCount.current += 1;
                    setShouldRender((prev) => !prev);
                } else {
                    clearInterval(interval);
                }
            }, 10);
        });
    }, []);


    return <h3 className="scroll-m-20 text-2xl tracking-tight text-center">
            Service status for {currentCount.current} companies!
        </h3>
}
