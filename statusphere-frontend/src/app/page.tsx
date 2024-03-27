"use client"

import {Search} from "@/components/Search";
import {Header} from "@/components/header";
import {Footer} from "@/components/footer";
import {Meteors} from "@/components/mainsite/Meteors";


function MainContent() {
    return (
        <div className={"flex justify-center w-full"}>
            <div className={"sm:w-[80vw] lg:w-[40vw]"}>
                <Search/>
            </div>
        </div>
    )
}

export default function Home() {
    return <>
        <Header/>
        <Meteors number={20}/>
        <MainContent/>
        <Footer/>
    </>
}
