"use client"

import {Search} from "@/components/Search";
import {Header} from "@/components/header";
import {Meteors} from "@/components/mainsite/Meteors";
import {CompanyCount} from "@/components/CompanyCount";
import Image from 'next/image'
import statusphere from "../../public/static/images/statusphere.png";
import {Button} from "@/components/ui/button";
import {useRouter} from "next/navigation";

function MainContent() {

    return (
        <div className={"flex justify-center h-full w-[100w] z-10 m-4"}>
            <div className={" h-full lg:w-[40vw] w-[100w]"}>
                <div className={" h-max flex justify-center"}>
                <div className={"flex justify-center w-[30vh] h-[30vh] relative"}>
                <Image
                    src={statusphere}
                    layout='fill'
                    objectFit='contain'
                    alt="Statusphere logo"
                />
                </div>
                </div>
                <CompanyCount/>
                <Search/>
            </div>
        </div>
    )
}

export default function Home() {
    return <>
        <Header/>
        <Meteors number={10}/>
        <MainContent/>
    </>
}
