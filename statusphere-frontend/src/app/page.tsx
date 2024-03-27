import {Header} from "@/components/header";
import {Footer} from "@/components/footer";
import {Meteors} from "@/components/mainsite/Meteors";

export default function Home() {
    return (
        <>
            <Header/>
            <Meteors number={20}/>
            <main>
                Status Aggregator WIP
            </main>
            <Footer/>
        </>

    );
}
