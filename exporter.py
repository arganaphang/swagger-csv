import sys
import json
from dataclasses import dataclass


@dataclass
class Row:
    endpoint: str
    method: str
    description: str
    summary: str
    consumes: str
    produces: str
    responses: any


def parse(path):
    rows = []
    with open(path, "r") as f:
        data = json.load(f)
        for endpoint in data["paths"]:
            for method in data["paths"][endpoint]:
                responsesStr = ""
                responses = data["paths"][endpoint][method]["responses"]
                for response in responses:
                    ref = responses[response]["schema"]["$ref"].rsplit("/", 1)[-1]
                    res = {
                        "status_code": response,
                        "description": responses[response]["description"],
                        "response": data["definitions"][ref],
                    }
                    responsesStr += json.dumps(res)
                rows.append(
                    Row(
                        endpoint=endpoint,
                        method=method,
                        description=data["paths"][endpoint][method]["description"],
                        summary=data["paths"][endpoint][method]["summary"],
                        consumes=", ".join(data["paths"][endpoint][method]["consumes"]),
                        produces=", ".join(data["paths"][endpoint][method]["produces"]),
                        responses=responsesStr,
                    )
                )
    with open(path + ".csv", "w") as f:
        f.write("endpoint;method;description;summary;consumes;produces;responses\n")
        for row in rows:
            f.write(
                "{};{};{};{};{};{};{}\n".format(
                    row.endpoint,
                    row.method,
                    row.description,
                    row.summary,
                    row.consumes,
                    row.produces,
                    row.responses,
                )
            )
    f.close()


if __name__ == "__main__":
    # Validate Args
    if len(sys.argv) < 2:
        print("Usage: {} [file containing PATH]").format(sys.argv[0])
        sys.exit(1)
    # Run Exporter
    parse(sys.argv[1])
