import {Header} from "@/components/header";
import {Footer} from "@/components/footer";
import {Meteors} from "@/components/mainsite/Meteors";

const Layout = ({children}) => {
    return (
        <>
            <Header/>
            <Meteors number={20}></Meteors>
            <div>
                <main>{children}</main>
            </div>
            <Footer/>
        </>
    );
};
export default Layout;