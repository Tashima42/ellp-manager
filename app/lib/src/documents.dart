import 'package:flutter/material.dart';

enum DocumentType { rg, cpf, enrollmentProof, other }

const documentTypes = {
  "RG": DocumentType.rg,
  "CPF": DocumentType.cpf,
  "Declaração de matrícula": DocumentType.enrollmentProof,
  "outros": DocumentType.other,
};

class RequestDocumentPage extends StatefulWidget {
  const RequestDocumentPage({super.key, required this.title});
  final String title;

  @override
  State<RequestDocumentPage> createState() => _RequestDocumentPageState();
}

class _RequestDocumentPageState extends State<RequestDocumentPage> {
  final _formKey = GlobalKey<FormState>();
  int? userID;
  String? name;
  DocumentType? type;
  String? keyWords;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Form(
        key: _formKey,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: <Widget>[
            DropdownButtonFormField(
              decoration: const InputDecoration(
                icon: Icon(Icons.assignment),
                labelText: "Tipo",
                hintText: "Selecione um tipo de documento",
              ),
              items: documentTypes.keys
                  .map<DropdownMenuItem<DocumentType>>((String documentType) {
                return DropdownMenuItem(
                  value: documentTypes[documentType],
                  child: Text(documentType),
                );
              }).toList(),
              onChanged: (DocumentType? value) => type = value,
              validator: (DocumentType? value) {
                if (value == null) return "Este campo é obrigatório";
                return null;
              },
            ),
            // update this to a searcheable field
            DropdownButtonFormField(
              decoration: const InputDecoration(
                icon: Icon(Icons.person),
                labelText: "Usuário que deve enviar",
                hintText: "Selecione um usuário",
              ),
              items: const [
                DropdownMenuItem(
                  value: 1,
                  child: Text("Joao Silva"),
                ),
                DropdownMenuItem(
                  value: 2,
                  child: Text("Brisa Dalila"),
                ),
              ],
              onChanged: (int? value) => userID = value,
              validator: (int? value) {
                if (value == null) return "Este campo é obrigatório";
                return null;
              },
            ),
            TextFormField(
              decoration: const InputDecoration(
                icon: Icon(Icons.description),
                labelText: "Nome do documento (opcional)",
                hintText: "atestado médico",
              ),
              onSaved: (String? value) => name = value,
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () => {},
        tooltip: 'Save',
        child: const Icon(Icons.save),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
