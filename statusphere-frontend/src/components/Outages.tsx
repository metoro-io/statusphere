import {Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";

interface OutagesProps {
    displayName: string;
    incidents: Incident[];
    statusPageUrl: string;
}

export function Outages(props: OutagesProps) {
    const convertToLink = (url: string, display: string) => {
        return <a
            style={{color: 'green'}}
            href={url}>{display}</a>;
    }

    return (
        <div>
            <Table>
                <TableCaption> Source:
                    {convertToLink(props.statusPageUrl, " Official " +  props.displayName + " status page")}
                </TableCaption>
                <TableHeader>
                    <TableRow>
                        <TableHead className="w-[100px]">Start Time (UTC)</TableHead>
                        <TableHead>End Time</TableHead>
                        <TableHead>Impact</TableHead>
                        <TableHead className="text-left">Duration</TableHead>
                        <TableHead className="text-left">Description</TableHead>
                        <TableHead className="text-left">Incident Page</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {props.incidents.map((incident) => (
                        <TableRow>
                            <TableCell>{convertSimpleDate(incident.StartTime)}</TableCell>
                            <TableCell>{incident.EndTime}</TableCell>
                            <TableCell>{incident.Impact}</TableCell>
                            <TableCell>{calculateDuration(incident.StartTime, incident.EndTime)}</TableCell>
                            <TableCell className="text-left">{incident.Description}</TableCell>
                            <TableCell className="text-left">{convertToLink(incident.DeepLink, "here")}</TableCell>

                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </div>
    );
}

const calculateDuration = (startTime: string, endTime: string): string => {
    // Parse the start and end times
    const start = new Date(startTime);
    const end = new Date(endTime);

    // Calculate the difference in milliseconds
    const difference = end.getTime() - start.getTime();

    // Convert milliseconds to hours, minutes, and seconds
    const hours = Math.floor(difference / (1000 * 60 * 60));
    const minutes = Math.floor((difference % (1000 * 60 * 60)) / (1000 * 60));
    const seconds = Math.floor((difference % (1000 * 60)) / 1000);

    // Build the duration string conditionally
    let durationParts = [];
    if (hours > 0) {
        const plural = hours === 1 ? "" : "s";
        durationParts.push(`${hours} hour${plural}`);
    }
    if (minutes > 0) {
        const plural = minutes === 1 ? "" : "s";
        durationParts.push(`${minutes} minute${plural}`);
    }
    if (seconds > 0 || durationParts.length === 0) {
        durationParts.push(`${seconds} seconds`);
    }
    return durationParts.join(", ");
};

const convertSimpleDate = (isoDateTime: string): string => {
    const date = new Date(isoDateTime);

    const year = date.getFullYear();
    // Months in JavaScript are 0-indexed, so add 1 for the correct month number
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}`;
};

