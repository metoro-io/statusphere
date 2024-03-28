import {Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {useEffect, useState} from "react";
import axios from "@/utils/axios";
import {StatusPage} from "@/model/StatusPage";
import {calculateDuration, convertToSimpleDate} from "@/utils/datetime";
import {ReadMore} from "@/components/ReadMore";

interface OutagesProps {
    statusPageDetails: StatusPage;
}

export function Outages(props: OutagesProps) {
    const [incidents, setIncidents] = useState<Incident[]>([]);

    useEffect(() => {
        const getIncidents = async () => {
            try {
                const response = await axios.get(
                    '/api/v1/incidents?statusPageUrl=' + props.statusPageDetails.url
                );
                setIncidents(response.data.incidents)
            } catch (err) {
                console.log(err);
            }
        };
        if (props.statusPageDetails.isIndexed) {
            getIncidents();
        }
    }, [props.statusPageDetails]);

    if (!props.statusPageDetails.isIndexed) {
        return <div>Incidents are not indexed for this status page,
            you can view the official status page in: <a
                href={props.statusPageDetails.url}> Official {props.statusPageDetails.name} status page</a>
        </div>
    }

    if (incidents === undefined || incidents.length === 0) {
        return <div>No incidents found</div>
    }

    return <>
        <h2 className="scroll-m-20 border-b pb-2 text-2xl font-semibold first:mt-0 text-center">
            Past {props.statusPageDetails.name} Incidents
        </h2>
    <Table className={"bg-white"}>
        <TableCaption> Source:
            <a href={props.statusPageDetails.url}> Official {props.statusPageDetails.name} status page</a>
        </TableCaption>
        <TableHeader>
            <TableRow>
                <TableHead>Start Time (UTC)</TableHead>
                <TableHead className={"max-w-[300px]"}>Incident Deep Link</TableHead>
                <TableHead>Impact</TableHead>
                <TableHead className="text-left">Duration</TableHead>
                <TableHead className="text-left">Description</TableHead>
            </TableRow>
        </TableHeader>
        <TableBody>
            {incidents.map((incident) => (
                <TableRow>
                    <TableCell>{convertToSimpleDate(incident.startTime)}</TableCell>
                    <TableCell><a href={incident.deepLink}> {incident.title} </a></TableCell>
                    <TableCell>{incident.impact}</TableCell>
                    <TableCell>{calculateDuration(incident.startTime, incident.endTime)}</TableCell>
                    <TableCell className="text-left">
                        <ReadMore text={incident.description}/>
                    </TableCell>
                    <TableCell className="text-left">
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
    </Table>
        </>;
}


