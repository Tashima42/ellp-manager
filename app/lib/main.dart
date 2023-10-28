import 'package:email_validator/email_validator.dart';
import 'package:flutter/material.dart';

enum Role { student, coordinator, instructor, secretary }

const roles = {
  "Estudante": Role.student,
  "Coordenador": Role.coordinator,
  "Instrutor": Role.instructor,
  "Secretário": Role.secretary,
};

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'ELLP Manager',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(
            seedColor: const Color.fromARGB(0, 0, 123, 204)),
        useMaterial3: true,
      ),
      home: const UserPage(title: 'Cadastro de usuário'),
    );
  }
}

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
  Role? role;
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
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.primary,
        title: Text(
          widget.title,
          style: const TextStyle(color: Colors.white),
        ),
      ),
      body: Padding(
        padding: const EdgeInsets.fromLTRB(15.0, 5.0, 15.0, 5.0),
        child: Form(
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
                items: roles.keys.map<DropdownMenuItem<Role>>((String role) {
                  return DropdownMenuItem(
                    value: roles[role],
                    child: Text(role),
                  );
                }).toList(),
                onChanged: (Role? value) {
                  role = value;
                },
                validator: (Role? value) {
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
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _printName,
        tooltip: 'Save',
        child: const Icon(Icons.save),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
