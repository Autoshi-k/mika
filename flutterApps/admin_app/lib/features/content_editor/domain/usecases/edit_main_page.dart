import 'package:flutter_apps/core/error/failures.dart';
import 'package:flutter_apps/features/content_editor/domain/repositories/main_page_editor_repository.dart';
import '../entities/main_page.dart';

class EditMainPage {
  final MainPageEditorRepository repository;
  EditMainPage(this.repository);

  Future<Failure?> saveTempChanges({
    required MainPage,
  }) async {
    return null;
  }
}