import 'package:dartz/dartz.dart';
import 'package:flutter_apps/features/content_editor/domain/entities/main_page.dart';
import 'package:flutter_apps/features/content_editor/domain/repositories/main_page_editor_repository.dart';
import 'package:flutter_apps/features/content_editor/domain/usecases/edit_main_page.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

class MockMainPageEditorRepository extends Mock implements MainPageEditorRepository {}

void main() {
  late EditMainPage usecase;
  late MockMainPageEditorRepository repository;

  setUp(() {
    repository = MockMainPageEditorRepository();
    usecase = EditMainPage(repository);
  });

  final tMainPage = MainPage(text: 'Hello', logo: 'child');
  
  test('should get null (as an error)', () async {
    when(repository.saveTempData(tMainPage)).thenAnswer((_) async => null);
    when(repository.getCurrentData()).thenAnswer((_) async => Right(tMainPage));

    final result = await usecase.saveTempChanges(MainPage: tMainPage);

    expect(result, null);
  });
}