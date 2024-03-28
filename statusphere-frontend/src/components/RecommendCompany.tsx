import {Card, CardHeader, CardTitle} from "@/components/ui/card";
interface RecommendCompanyProps {
    input: string;
}

export function RecommendCompany(props: RecommendCompanyProps) {
    return <Card className={"bg-white"}>
                <CardHeader className={"items-left"}>
                    <CardTitle>
                        <h3 className="scroll-m-20 text-xl font-semibold tracking-tight">
                            We currently don&apos;t have status updates for {props.input}
                        </h3>
                        <p className="leading-7 [&:not(:first-child)]:mt-6">
                            To add a company status page, please make a pull request in <a href={"https://github.com/metoro-io/statusphere"}>Statusphere repository</a>.
                        </p>
                    </CardTitle>
                </CardHeader>
            </Card>
}