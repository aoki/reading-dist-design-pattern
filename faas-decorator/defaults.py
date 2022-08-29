def handler(context):
    obj = context.json
    if obj.get("name", None) is None:
        obj["name"] = "John Doe"
    if obj.get("color", None) is None:
        obj["color"] = "blue"
    return obj
