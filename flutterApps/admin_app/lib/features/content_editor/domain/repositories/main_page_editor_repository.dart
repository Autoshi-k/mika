import 'package:dartz/dartz.dart';

import '../../../../core/error/failures.dart';
import '../entities/main_page.dart';

abstract class MainPageEditorRepository {
  Future<Either<Failure, MainPage>> getCurrentData();
  Future<Failure?> saveTempData(MainPage mainPage);
}