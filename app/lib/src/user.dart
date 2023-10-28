import 'package:flutter/material.dart';
import 'package:email_validator/email_validator.dart';

enum UserRole { student, coordinator, instructor, secretary }

const userRoles = {
  "Estudante": UserRole.student,
  "Coordenador": UserRole.coordinator,
  "Instrutor": UserRole.instructor,
  "Secretário": UserRole.secretary,
};

class UserPage extends StatefulWidget {
  const UserPage({super.key, required this.title});
  final String title;

  @override
  State<UserPage> createState() => _UserPageState();
}

class _UserPageState extends State<UserPage> {
  final _formKey = GlobalKey<FormState>();
  String? name;
  String? email;
  UserRole? role;
  String? password;
  String? address;

  void _printName() {
    if (!_formKey.currentState!.validate()) return;
    _formKey.currentState!.save();
    ScaffoldMessenger.of(context).showSnackBar(
      const SnackBar(content: Text('Salvando dados')),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Form(
        key: _formKey,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: <Widget>[
            TextFormField(
                decoration: const InputDecoration(
                  icon: Icon(Icons.person),
                  labelText: "Nome",
                  hintText: "João Silva",
                ),
                onSaved: (String? value) => name = value,
                validator: (String? value) {
                  if (value == null || value.isEmpty) {
                    return "Este campo é obrigatório";
                  }
                  return null;
                }),
            TextFormField(
                decoration: const InputDecoration(
                  icon: Icon(Icons.email),
                  labelText: "Email",
                  hintText: "membro@alunos.utfpr.edu.br",
                ),
                keyboardType: TextInputType.emailAddress,
                onSaved: (String? value) => email = value,
                validator: (String? value) {
                  if (value == null || value.isEmpty) {
                    return "Este campo é obrigatório";
                  }
                  if (!EmailValidator.validate(value)) {
                    return "Este email não é válido";
                  }
                  return null;
                }),
            DropdownButtonFormField(
              decoration: const InputDecoration(
                icon: Icon(Icons.recent_actors),
                labelText: "Função",
                hintText: "Selecione função do usuário",
              ),
              items: userRoles.keys.map<DropdownMenuItem<UserRole>>((String role) {
                return DropdownMenuItem(
                  value: userRoles[role],
                  child: Text(role),
                );
              }).toList(),
              onChanged: (UserRole? value) {
                role = value;
              },
              validator: (UserRole? value) {
                if (value == null) {
                  return "Este campo é obrigatório";
                }
                return null;
              },
            ),
            TextFormField(
                decoration: const InputDecoration(
                  icon: Icon(Icons.password),
                  labelText: "Senha",
                ),
                obscureText: true,
                onSaved: (String? value) => password = value,
                validator: (String? value) {
                  if (value == null || value.isEmpty) {
                    return "Este campo é obrigatório";
                  }
                  if (value.length < 5) {
                    return "A senha precisa ter ao menos 5 caracteres";
                  }
                  return null;
                }),
            TextFormField(
                decoration: const InputDecoration(
                  icon: Icon(Icons.house),
                  labelText: "Endereço",
                ),
                keyboardType: TextInputType.streetAddress,
                onSaved: (String? value) => address = value,
                validator: (String? value) {
                  if (value == null || value.isEmpty) {
                    return "Este campo é obrigatório";
                  }
                  return null;
                }),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _printName,
        tooltip: 'Save',
        child: const Icon(Icons.save),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
