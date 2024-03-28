"use client"

import {Search} from "@/components/Search";
import {Header} from "@/components/header";
import {Footer} from "@/components/footer";
import {Meteors} from "@/components/mainsite/Meteors";
import {CompanyCount} from "@/components/CompanyCount";
import Image from 'next/image'
import statusphere from "../../public/static/images/statusphere.png";



function MainContent() {
    return (
        <div className={"flex justify-center w-full z-10 m-4"}>
            <div className={"sm:w-[80vw] lg:w-[40vw]"}>
                <Image
                    src={statusphere}
                    width={500}
                    height={500}
                    alt="Picture of the author"
                />
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
        <Footer/>
    </>
}
