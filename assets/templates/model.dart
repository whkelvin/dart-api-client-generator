import 'package:json_annotation/json_annotation.dart';

part '{{.FileName}}.g.dart';

@JsonSerializable()
class {{.ClassName}} {
  final String firstName;

  {{.ClassName}}({required this.firstName});

  factory {{.ClassName}}.fromJson(Map<String, dynamic> json) => _${{.ClassName}}FromJson(json);
  Map<String, dynamic> toJson() => _${{.ClassName}}ToJson(this);
}
