import {StatusPage} from "@/model/StatusPage";
import axios from "@/utils/axios";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {NavLink} from "@/components/mainsite/NavLink";

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


export default function Directory({companyList}: DirectoryProps) {
    return <div className={"m-4"}>
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
                    <TableRow>
                        <TableCell>{company.name}</TableCell>
                        <TableCell>
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