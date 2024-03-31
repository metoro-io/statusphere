import {StatusPage} from "@/model/StatusPage";
import axios from "@/utils/axios";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {NavLink} from "@/components/mainsite/NavLink";
import {Metadata} from "next";
import {Props} from "next/script";
import Head from "next/head";

interface DirectoryProps {
    companyList: StatusPage[]
}

export async function getServerSideProps() {
    const response = await axios.get('/api/v1/statusPages')
    return {
        props: {
            companyList: response.data.statusPages
        }
    }
}


export async function generateMetadata() {
    return {
        title: `Statusphere - Directory`,
    }
}

export default function Directory({companyList}: DirectoryProps) {
    return <div className={"m-4"}>
        <Head>
            <title>Statusphere - Directory</title>
            <meta name="description" content="Statusphere Company Directory"/>
            <meta name="keywords" content="status, statusphere, statuspage, company, directory"/>
        </Head>
        <h3 className="scroll-m-20 text-2xl font-semibold tracking-tight">
            Status Page List
        </h3>
        <Table className={"bg-white"}>
            <TableHeader>
                <TableRow>
                    <TableHead>Company</TableHead>
                    <TableHead>Status Page</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {companyList.map((company) => (
                    <TableRow key={company.name + ""}>
                        <TableCell key={company.name + "-name"}>{company.name}</TableCell>
                        <TableCell key={company.name + "-link"}>
                            <NavLink
                                href={"/status/" + company.name}>See {company.name} status
                                here</NavLink>
                        </TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    </div>
}