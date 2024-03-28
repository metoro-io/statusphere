import {Header} from "@/components/header";
import {Footer} from "@/components/footer";
import {Meteors} from "@/components/mainsite/Meteors";

const Layout = ({children}) => {
    return (
        <div className={"w-[100vw]"}>
            <Header/>

            <div className={"z-10"}>
                {children}
            </div>
        </div>
    );
};
export default Layout;