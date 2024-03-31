import {Header} from "@/components/header";
import {ReactNode} from "react";
import Head from "next/head";

interface LayoutProps {
    children: ReactNode;
}
const Layout = (props: LayoutProps) => {
    return (
        <div className={"w-[100vw]"}>
            <Head>
                <html lang={"en"}/>
            </Head>
            <Header/>
            <div className={"z-10"}>
                {props.children}
            </div>
        </div>
    );
};
export default Layout;