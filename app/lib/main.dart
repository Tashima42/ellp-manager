import 'package:app/src/documents.dart';
import 'package:app/src/user.dart';
import 'package:flutter/material.dart';

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
      home: Scaffold(
          appBar: AppBar(
            title: const Text("Cadastro de usuário"),
            centerTitle: true,
            backgroundColor: const Color.fromARGB(133, 9, 108, 174),
            toolbarHeight: 40.0,
          ),
          bottomNavigationBar:
              BottomNavigationBar(currentIndex: 1, items: const [
            BottomNavigationBarItem(
                icon: Icon(Icons.calendar_month), label: "Eventos"),
            BottomNavigationBarItem(
              icon: Icon(Icons.person),
              label: "Usuários",
            ),
            BottomNavigationBarItem(
                icon: Icon(Icons.folder_open), label: "Documentos")
          ]),
          body: const Padding(
            padding: EdgeInsets.fromLTRB(15.0, 5.0, 15.0, 5.0),
            // child: UserPage(title: 'Cadastro de usuário'),
            child: RequestDocumentPage(title: 'Requisitar Documento'),
          )),
    );
  }
}
