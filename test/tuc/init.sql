-- initialize database to do December 2013 tucasa rentroll

USE rentroll

-- define the property
INSERT INTO property (Name,Address,Address2,City,State,PostalCode,Country,Phone,DefaultOccupancyType,ParkingPermitInUse) VALUES
	("Tucasa Townhomes","1635 Tucasa Drive","","Irving","TX","75061-3179","USA","469-284-9506",3,0);

-- define unit types for TUC
INSERT INTO unittypes (PRID,Designation,Description) VALUES
	(1,"Studio/1","385 sq ft"),
	(1,"1/1 Townhome","726 sq ft"),
	(1,"2/2 Townhome","770 sq ft"),
	(1,"1/3 Townhome","1,123 sq ft");

-- define unit specialties
INSERT INTO unitspecialtytypes (PRID,Name,Fee,Description) VALUES
	(1,"Lake View",0.0,"Overlooks the lake"),
	(1,"Courtyard View",0.0,"Rear windows view the courtyard"),
	(1,"Washer Dryer Connections",0.0,"Connections for your washer/dryer"),
	(1,"Washer Dryer provided",0.0,"Maytag washer/dryer provided"),
	(1,"Fireplace",0.0,"Wood burning, gas fireplace");

-- define the assessments for Tucasa
INSERT INTO propertyassessments (PRID,ASMID) VALUES
	(1, 1),  -- Vacancy
	(1, 2),  -- BadDebtExpense
	(1, 3),  -- LossToLease
	(1, 4),  -- AssociateConcession
	(1, 5),  -- ResidentConcession
	(1, 6),  -- OfflineUnit
	(1, 7),  -- OtherOffset
	(1, 8),  -- Rent
	(1, 9),  -- SecurityDeposit
	(1,10),  -- SecurityDepositForfeiture
	(1,11),  -- ApplicationFees
	(1,12),  -- LandlordLienSales
	(1,13),  -- PetFees
	(1,14),  -- EvictionFees
	(1,15),  -- ElectricReimbursement
	(1,16),  -- ElectricOverage
	(1,17),  -- WaterReimbursement
	(1,18),  -- WaterOverage
	(1,19),  -- TrashFee
	(1,20),  -- UtilityFine
	(1,21),  -- NSFFee
	(1,22),  -- MaintenanceFee
	(1,23),  -- Fines
	(1,24),  -- MonthToMonthFee
	(1,25),  -- CancelationFees
	(1,26),  -- HousekeepingFee
	(1,27),  -- ExtraPersonCharge
	(1,28),  -- FurnitureRental
	(1,29),  -- PlatinumServiceFee
	(1,33),  -- TOTTax
	(1,34),  -- SpecialEventFees
	(1,37),  -- VendingIncome
	(1,41),  -- WashNFoldIncome
	(1,42),  -- SpaSales
	(1,43),  -- FitnessCenter
	(1,44);  -- CashOverShort

-- define the building for TUC
INSERT INTO buildings (Address,Address2,City,State,PostalCode,Country) VALUES
	("1635 Tucasa Drive","","Irving","TX","75061-3179","USA");

-- define the units
INSERT INTO units (PRID,BLDGID,UTID,Name,DefaultOccType,OccType,ScheduledRent,Assignment) VALUES
	(1,1,1,"101",2,2, 1100.00, 1),
	(1,1,2,"102",2,2, 805.00, 1),
	(1,1,3,"103",2,2, 1060.0, 1),
	(1,1,4,"104",2,2, 550.00, 1),
	(1,1,5,"105",2,2, 0.00, 1),
	(1,1,6,"106",2,2, 995.00, 1),
	(1,1,7,"107",2,2, 1900.00, 1),
	(1,1,8,"108",2,2, 1085.00, 1),
	(1,1,9,"109",2,2, 1849.00, 1),
	(1,1,10,"110",2,2, 845.00, 1),
	(1,1,11,"111",2,2, 0.00, 1),
	(1,1,12,"112",2,2, 905.00, 1),
	(1,1,13,"113",2,2, 0.00, 1),
	(1,1,14,"114",2,2, 1690.00, 1),
	(1,1,15,"115",2,2, 1135.00, 1),
	(1,1,16,"116",2,2, 995.00, 1),
	(1,1,17,"117",2,2, 1890.00, 1),
	(1,1,18,"118",2,2, 975.00, 1),
	(1,1,19,"119",2,2, 700.00, 1),
	(1,1,20,"120",2,2, 0.00, 1),
	(1,1,21,"121",2,2, 935.00, 1),
	(1,1,22,"122",2,2, 0.00, 1),
	(1,1,23,"123",2,2, 725.00, 1),
	(1,1,24,"124",2,2, 830.00, 1),
	(1,1,25,"125",2,2, 0.00, 1),
	(1,1,26,"126",2,2, 745.00, 1),
	(1,1,27,"127",2,2, 1340.00, 1),
	(1,1,28,"128",2,2, 625.00, 1),
	(1,1,29,"129",2,2, 0.00, 1),
	(1,1,30,"130",2,2, 795.00, 1),
	(1,1,31,"131",2,2, 0.00, 1),
	(1,1,32,"132",2,2, 695.00, 1),
	(1,1,33,"133",2,2, 795.00, 1),
	(1,1,34,"134",2,2, 0.00, 1),
	(1,1,35,"135",2,2, 605.00, 1),
	(1,1,36,"136",2,2, 1135.00, 1),
	(1,1,37,"137",2,2, 945.00, 1),
	(1,1,38,"138",2,2, 885.00, 1),
	(1,1,39,"139",2,2, 365.00, 1),
	(1,1,40,"140",2,2, 805.00, 1),
	(1,1,41,"141",2,2, 0.00, 1),
	(1,1,42,"142",2,2, 1910.00, 1),
	(1,1,43,"143",2,2, 1070.00, 1),
	(1,1,44,"144",2,2, 1790.00, 1),
	(1,1,45,"145",2,2, 1080.00, 1),
	(1,1,46,"146",2,2, 975.00, 1),
	(1,1,47,"147",2,2, 0.00, 1),
	(1,1,48,"148",2,2,  35.00, 1),
	(1,1,49,"149",2,2, 740.00, 1),
	(1,1,50,"150",2,2, 550.00, 1),
	(1,1,51,"151",2,2, 1082.00, 1),
	(1,1,52,"152",2,2, 1050.00, 1),
	(1,1,53,"153",2,2, 550.00, 1),
	(1,1,54,"154",2,2, 1300.00, 1),
	(1,1,55,"155",2,2, 845.00, 1),
	(1,1,56,"158",2,2, 850.00, 1),
	(1,1,57,"159",2,2, 795.00, 1),
	(1,1,58,"160",2,2, 695.00, 1),
	(1,1,59,"161",2,2, 780.00, 1),
	(1,1,60,"162",2,2, 730.00, 1),
	(1,1,61,"163",2,2, 0.00, 1),
	(1,1,62,"164",2,2, 710.00, 1),
	(1,1,63,"167",2,2, 0.00, 1),
	(1,1,64,"168",2,2, 0.00, 1),
	(1,1,65,"201",2,2, 1100.00, 1),
	(1,1,66,"202",2,2, 2100.00, 1),
	(1,1,67,"203",2,2, 0.00, 1),
	(1,1,68,"204",2,2, 0.00, 1),
	(1,1,69,"205",2,2, 685.00, 1),
	(1,1,70,"206",2,2, 765.00, 1),
	(1,1,71,"207",2,2, 994.00, 1),
	(1,1,72,"208",2,2, 0.00, 1),
	(1,1,73,"209",2,2, 0.00, 1),
	(1,1,74,"210",2,2, 0.00, 1),
	(1,1,75,"211",2,2,  100.00, 1),
	(1,1,76,"212",2,2, 1390.00, 1),
	(1,1,77,"213",2,2, 695.00, 1),
	(1,1,78,"214",2,2, 800.00, 1),
	(1,1,79,"215",2,2, 46.00, 1),
	(1,1,80,"216",2,2, 1990.00, 1),
	(1,1,81,"217",2,2, 930.00, 1),
	(1,1,82,"218",2,2, 2260.00, 1),
	(1,1,83,"219",2,2, 0.00, 1),
	(1,1,84,"220",2,2, 930.00, 1),
	(1,1,85,"221",2,2, 0.00, 1),
	(1,1,86,"222",2,2, 880.00, 1),
	(1,1,87,"223",2,2, 291.00, 1),
	(1,1,88,"224",2,2, 997.00, 1),
	(1,1,89,"225",2,2, 885.00, 1),
	(1,1,90,"226",2,2, 145.00, 1),
	(1,1,91,"227",2,2, 300.00, 1),
	(1,1,92,"228",2,2, 840.00, 1),
	(1,1,93,"229",2,2, 865.00, 1),
	(1,1,94,"230",2,2, 1140.00, 1),
	(1,1,95,"231",2,2, 905.00, 1),
	(1,1,96,"232",2,2, 895.00, 1),
	(1,1,97,"233",2,2, 2037.00, 1),
	(1,1,98,"234",2,2, 0.00, 1),
	(1,1,99,"235",2,2, 975.00, 1),
	(1,1,100,"236",2,2, 795.00, 1),
	(1,1,101,"237",2,2, 2060.00, 1),
	(1,1,102,"238",2,2, 1990.00, 1),
	(1,1,103,"239",2,2, 261.00, 1),
	(1,1,104,"240",2,2, 0.00, 1),
	(1,1,105,"241",2,2, 945.00, 1),
	(1,1,106,"242",2,2, 890.00, 1),
	(1,1,107,"243",2,2, 1030.00, 1),
	(1,1,108,"244",2,2, 980.00, 1),
	(1,1,109,"245",2,2, 1500.00, 1),
	(1,1,110,"246",2,2, 937.00, 1),
	(1,1,111,"247",2,2, 1740.00, 1),
	(1,1,112,"248",2,2, 855.00, 1),
	(1,1,113,"249",2,2, 905.00, 1),
	(1,1,114,"250",2,2, 895.00, 1),
	(1,1,115,"251",2,2, 855.00, 1),
	(1,1,116,"252",2,2, 1440.00, 1),
	(1,1,117,"253",2,2, 720.00, 1),
	(1,1,118,"254",2,2, 725.00, 1),
	(1,1,119,"255",2,2, 275.00, 1),
	(1,1,120,"256",2,2, 805.00, 1),
	(1,1,121,"257",2,2, 1045.00, 1),
	(1,1,122,"258",2,2, 1015.00, 1),
	(1,1,123,"259",2,2, 509.00, 1),
	(1,1,124,"260",2,2, 695.00, 1),
	(1,1,125,"261",2,2, 835.00, 1),
	(1,1,126,"262",2,2, 0.00, 1),
	(1,1,127,"263",2,2, 1700.00, 1),
	(1,1,128,"264",2,2, 1251.00, 1);

-- Tenants


-- Macias
-- Morua
-- Garcia
-- Duran
-- Canas
-- Caballero_Valiente
-- Fonseca
-- Hernanadez
-- Martinez
-- Romero
-- Ragsdale_Anthony
-- Salinas
-- Gonzalez_Ibarra
-- Polite
-- Frazier
-- Palacios
-- Perez
-- Ortiz
-- Reyes
-- Jeri
-- Castaneda
-- Ramirez
-- Salazar
-- Casillas
-- remodeled_unit
-- Guiterrez
-- Rivera
-- Marrufo
-- Hernandez
-- Ali
-- Aguilar
-- McDaniel
-- Reed
-- Avalos_Barrera
-- remodeled_Jaramillo
-- Zapata
-- Palacios
-- Ramirez
-- remodeled_Gomez
-- Barrera
-- Moreno
-- Hernanadez_Lopez
-- Flores
-- Juarez
-- Rivera
-- Flores
-- Ramirez
-- Tinoco
-- Gonzales
-- Sexton
-- Mariscal
-- Brown
-- Bable
-- Hernandez
-- Martinez
-- Moradel
-- Campos
-- Rodriquez
-- Flores
-- Garza
-- Gonzalez_Ramirez
-- Guerrero
-- Vidales
-- Ipina
-- ALI
-- Medelin
-- Villanueva
-- Vargas_Morales
-- Garcia
-- Rios
-- vacant_Benitez
-- Moraira
-- Santos
-- Vacant_Marquez
-- Guiterrez_Ortiz
-- Isa
-- Lara
-- Martinez
-- Moreno
-- Vargas
-- Carillo
-- Garcia
-- Santillana
-- Ventura
-- Valenciano
-- Rodriquez
-- Jaramillo
-- Martinez
-- Vasquez
-- Romano_Rodriguez
-- remodeledunit_Gonzalez
-- Muy
-- Benitez
-- Senobo
-- Gutierrez
-- Carrazo
-- Vazzquez
-- Palacios
-- Luna
-- 0
-- Adeoqun
-- Martinez
-- Remodeled_Unit_Lopez_Garcia
-- Cisneros
-- Fuents
-- Barbosa
-- Vazquez
-- Castillo
-- Lopez
-- Sorto
-- Moreno
-- Molina
-- Ojeda
-- Gonzalez
-- Orozco
-- Partin
-- lopez
-- Alvarez
-- Bejega_Kidiga
-- Beserra
-- Graciano
-- Garcia
-- Lopez
-- Pabafox
-- McDonald
-- Romano
-- Cruz
-- DeLaGarza_Andino
