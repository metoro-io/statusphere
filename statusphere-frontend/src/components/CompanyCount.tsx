import {useEffect, useRef, useState} from "react";
import axios from "@/utils/axios";
import {Button} from "@/components/ui/button";
import {useRouter} from "next/navigation";


export function CompanyCount() {
    const [count, setCount] = useState(0);
    const currentCount = useRef(0);
    const [shouldRender, setShouldRender] = useState(false);
    const router = useRouter();


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
            }, 5);
        });
    }, []);

    return <div className="flex justify-center">
        <div className="text-2xl tracking-tight">
        Service status for
            <Button className={"shadow-white text-xl p-1"} onClick={() => {
                router.push('/directory')
            }}><b className="text-[#00243c]">{currentCount.current}</b></Button>companies!
        </div>
    </div>

}
