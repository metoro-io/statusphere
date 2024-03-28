import {Header} from "@/components/header";
import {ReactNode} from "react";

interface LayoutProps {
    children: ReactNode;
}
const Layout = (props: LayoutProps) => {
    return (
        <div className={"w-[100vw]"}>
            <Header/>
            <div className={"z-10"}>
                {props.children}
            </div>
        </div>
    );
};
export default Layout;